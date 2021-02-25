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

mydb = mysql.connector.connect(
    host="match-mariadb-crossfitlov",
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




no_found = {'matche_?' : "No"}
Yes_found = {'matche_?' : "Yes"}

# La requête POST - Retourne un JSON en fonction des informations qui sont spécifiées
@app.route('/api/matches/', methods = ['POST'])
def matches():
    code, description = checkToken(request.headers.get('X-Authorization'))
    if not code == 200 : #si le token n'a pas plus être validé, ou alors si le token n'est pas valide
        abort(code, description=description)

    query = request.get_json()
        
    for _, dico in enumerate(query):
        id_1 = dico['id_1']
        id_2 = dico['id_2']
        
    mycursor = mydb.cursor(buffered=True, dictionary=True)
    mycursor.execute("SELECT * FROM matches WHERE usr_id_1 = " + id_1 + " AND usr_id_2 = " + id_2 )
    myresult = mycursor.fetchall()
    
    if not myresult :
        print("ERROR : NO MATCHES FOR -------------- usr_id_1 = " + id_1 + ", usr_id_2 = " + id_2 + " --------------")
        return jsonify(No_found) 
    else :
        return jsonify(Yes_found)
    
@app.route('/api/matches/', methods = ['PUT'])
def swipes() :
    code, description = checkToken(request.headers.get('X-Authorization'))
    if not code == 200 : #si le token n'a pas plus être validé, ou alors si le token n'est pas valide
        abort(code, description=description)

    query = request.get_json()
    
    for _, dico in enumerate(query):
        r_usr_id = dico['usr_id'] # Alexis Ren
        r_swipe_id = dico['swipe_with'] # Hugo Da Costa Pina
        
    insert_swipe(r_usr_id, r_swipe_id)
    mycursor = mydb.cursor(buffered=True, dictionary=True)
    mycursor.execute("SELECT * FROM swipes WHERE usr_id =" + r_swipe_id + " AND swipe_with =" +  r_usr_id + "")
    myresult = mycursor.fetchall() 
    
    if not myresult :
        print("ERROR : -------------- La requet sql est vide --------------")
        return ""
    else :
        insert_matches(r_usr_id, r_swipe_id)
        return ""
        
def insert_matches(id_1, id_2):
    sql = "INSERT INTO matches(usr_id_1, usr_id_2) VALUES (%s, %s)"
    val = (id_1, id_2)
    mycursor.execute(sql, val)
    mydb.commit()
    print("MATCHES --------------------- record inserted.--------------------------")
    
def insert_swipe(id_1, id_2):
    
    sql = "INSERT INTO swipes(usr_id, swipe_with) VALUES (%s, %s)"
    val = (id_1, id_2)
    mycursor.execute(sql, val)
    mydb.commit()
    print("SWIPE --------------------- record inserted.--------------------------")
    
if __name__ == "__main__":
    app.run(host='0.0.0.0', port=8080)