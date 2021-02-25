from flask import Flask, render_template, jsonify, request, abort
from flask_basicauth import BasicAuth
import json
import mysql.connector
import requests

app = Flask(__name__)

##################
### brique auth
##################
URL_AUTHAPI = 'http://auth-api-crossfitlov:8000/v1'
LOGIN_BASICAUTH_AUTHAPI = 'admin'
PASSWORD_BASICAUTH_AUTHAPI = 'admin123'

#############################################################################
### BASIC AUTH
#############################################################################
app.config['BASIC_AUTH_USERNAME'] = 'admin'
app.config['BASIC_AUTH_PASSWORD'] = 'admin123'
basic_auth = BasicAuth(app)
app.config['BASIC_AUTH_FORCE'] = True


#############################################################################
### PARTIE REQUETES PUT ET POST
#############################################################################

# YYYY-MM-DD hh:mm:ss

mydb = mysql.connector.connect(
  host="message-mariadb-crossfitlov",
  user="root",
  password="admin",
  database="CL"
)

mycursor = mydb.cursor()

#error handler for not acceptable requests
@app.errorhandler(406)
def resource_not_acceptable(e):
    return jsonify(error=str(e)), 406

#error handler for service unavailable
@app.errorhandler(503)
def resource_unavailable(e):
    return jsonify(error=str(e)), 503

#error handler for internal server error
@app.errorhandler(500)
def internal_server_error(e):
    return jsonify(error=str(e)), 500

#error handler for bad request
@app.errorhandler(400)
def badrequest_error(e):
    return jsonify(error=str(e)), 400

def checkToken(authorizationHeader) :

    #récupération du token
    prefix = 'Bearer '
    if not authorizationHeader:
        return 406, "No token provided"

    if not authorizationHeader.startswith(prefix):
        return 406, "Invalid authorization token format"
    
    token = authorizationHeader[len(prefix):]
    r = None
    try:
        r = requests.post(URL_AUTHAPI+'/check', auth=(LOGIN_BASICAUTH_AUTHAPI,PASSWORD_BASICAUTH_AUTHAPI), json = {'token':token}, timeout = 5)
    except (requests.exceptions.ConnectionError, requests.exceptions.Timeout) as e:
        return 503, "Service unavailable"
    
    if r.status_code == 401:
        return 503, "Service unavailable"

    if not r.status_code == 200:
        return r.status_code, "Invalid token"
	
    return 200, ""




# La requête POST - Retourne un JSON en fonction des informations qui sont spécifiées
@app.route('/api/message/', methods = ['POST'])
def message():
    code, description = checkToken(request.headers.get('X-Authorization'))
    if not code == 200 : #si le token n'a pas plus être validé, ou alors si le token n'est pas valide
        abort(code, description=description)

    query = request.get_json()
        
    for _, dico in enumerate(query):
        question = dico['room_id']
        
    mycursor = mydb.cursor(dictionary=True)
    mycursor.execute("SELECT * FROM messages WHERE room_id =" + question + " ORDER BY date_time DESC LIMIT 20")
    myresult = mycursor.fetchall()
    
    return jsonify(myresult)

# La requête PUT - Retourne un JSON avec l'ajout des nouvelles informations
@app.route('/api/message/sql', methods = ['PUT'])
def publi_message():
    
    code, description = checkToken(request.headers.get('X-Authorization'))
    if not code == 200 : #si le token n'a pas plus être validé, ou alors si le token n'est pas valide
        abort(code, description=description)

    query = request.get_json()
    
    for _,dico in enumerate(query):
        if dico['expediteur']:
            e_id = dico['expediteur']
            insert_user(e_id)
        if dico['room_id']:
            r_id = dico['room_id']
            room_id(r_id)
        if dico['message']:
            msg_text = dico['message']
            r_id = dico['room_id']
            e_id = dico['expediteur']
            date = dico['date']
            message(msg_text, r_id, e_id, date)
            
    return ""

def insert_user(id):
    
    sql = "INSERT IGNORE INTO users (id_usr) VALUES (" + id + ")" 
    mycursor.execute(sql)
    mydb.commit()
    print("useur_id --------------------- record inserted.--------------------------")

def room_id(id):
    
    sql = "INSERT IGNORE INTO rooms (room_id) VALUES (" + id + ")" 
    mycursor.execute(sql)
    mydb.commit()
    print("room_id  --------------------- record inserted.--------------------------")
    
def message(msg_text, r_id, e_id, date):
    
    sql = "INSERT INTO messages (msg_text, room_id, id_usr , date_time) VALUES (%s, %s, %s, %s)"
    val = (msg_text, r_id, e_id, date)
    mycursor.execute(sql, val)
    mydb.commit()
    print("msg_text --------------------- record inserted.--------------------------")

#############################################################################
### MAIN
#############################################################################

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=8080)