from flask import Flask, render_template, jsonify, request, abort
import json
import mysql.connector
from flask_basicauth import BasicAuth
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
  host="abonnement-mariadb-crossfitlov",
  user="root",
  password="admin",
  database="CL"
)

n_found = {'Abonnement' : "n"}
y_found = {'Abonnement' : "y"}


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
@app.route('/api/abonnement/', methods = ['POST'])
def abonnement():
    code, description = checkToken(request.headers.get('X-Authorization'))
    if not code == 200 : #si le token n'a pas plus être validé, ou alors si le token n'est pas valide
        abort(code, description=description)

    query = request.get_json()
        
    for _, dico in enumerate(query):
        question = dico['id_usr']
        
    mycursor = mydb.cursor(dictionary=True)
    mycursor.execute("SELECT * FROM abonnement WHERE id_usr = " + question )
    myresult = mycursor.fetchall()
    
    if not myresult :
        print("ERROR : NO ABBONEMENT FOR -------------- usr_id_1 = " + question)
        return jsonify(n_found) 
    else :
        return jsonify(y_found)
    
    return jsonify(myresult)
    
# La requête PUT - Retourne un JSON avec l'ajout des nouvelles informations
@app.route('/api/abonnement/', methods = ['PUT'])
def publi_abonnement():
    code, description = checkToken(request.headers.get('X-Authorization'))
    if not code == 200 : #si le token n'a pas plus être validé, ou alors si le token n'est pas valide
        abort(code, description=description)

    
    query = request.get_json()
    
    for _, dico in enumerate(query):
        id_usr = dico['id_usr'] 
        y_or_not = dico['y_or_n'] 
        
    insert_user(id_usr, y_or_not)
        
    return ""
    

def insert_user(id_usr, y_or_n):
    
    sql = "INSERT INTO abonnement(id_usr, y_or_n) VALUES (%s, %s)"
    val = (id_usr, y_or_n)
    mycursor.execute(sql, val)
    mydb.commit()
    print("usr_id --------------------- record inserted.--------------------------")


#############################################################################
### MAIN
#############################################################################

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=8080)