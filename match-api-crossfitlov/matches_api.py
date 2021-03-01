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

##################
### brique auth
##################
URL_USERSAPI = 'http://users-api-crossfitlov:8000/v1'
LOGIN_BASICAUTH_USERSAPI = 'admin'
PASSWORD_BASICAUTH_USERSAPI = 'admin123'

####################
### CORS middleware
####################
def CORS(response):
    response.headers['Access-Control-Allow-Origin'] = '*'
    if request.method == 'OPTIONS':
        response.headers['Access-Control-Allow-Methods'] = 'DELETE, GET, POST, PUT'
        headers = request.headers.get('Access-Control-Request-Headers')
        if headers:
            response.headers['Access-Control-Allow-Headers'] = headers
        response.status_code = 200
    return response
app.after_request(CORS)

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

#error handler for not found uri requests
@app.errorhandler(404)
def resource_not_found(e):
    return jsonify(error=str(e)), 404

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

def getUsers() :

    r = None
    try:
        r = requests.post(URL_USERSAPI+'/users/get', auth=(LOGIN_BASICAUTH_USERSAPI,PASSWORD_BASICAUTH_USERSAPI), timeout = 5)
    except (requests.exceptions.ConnectionError, requests.exceptions.Timeout) as e:
        return None, 503, "Service unavailable"
    if r.status_code != 200:
        return None, 503, "Brique users a retourné le code "+str(r.status_code)
    return r.json(), 200, ""


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



@app.route('/api/show/', methods = ['POST'])
def show() :
    code, description = checkToken(request.headers.get('X-Authorization'))
    if not code == 200 : #si le token n'a pas plus être validé, ou alors si le token n'est pas valide
        abort(code, description=description)

    query = request.get_json()

    mec_courant = query['usr_id'] # Alexis Ren


    #récuperation de la liste des utilisateurs
    userInfosList, code, description = getUsers()
    if not code == 200 : #si le token n'a pas plus être validé, ou alors si le token n'est pas valide
        abort(500, description=description)

    
    #[
    #    {
    #        "crossfitid"
    #        "biography"
    #        "job"
    #    }
    #]

    mycursor = mydb.cursor(buffered=True, dictionary=True)
    mycursor.execute("SELECT * FROM swipes WHERE usr_id = " + mec_courant)
    sqlresult = mycursor.fetchall()

    # sqlresult = ({usr_id_1:Hugo,usr_id_2:Alexisren},  {usr_id_1:Hugo,usr_id_2:Maxime}, {usr_id_1:Hugo,usr_id_2:Clara})

    idBDDSwipeList = []
    for _, v in enumerate(sqlresult):
        idBDDSwipeList.append(v['swipe_with'])

    # userInfosList = ({info:Hugo, info:Alexisren, crossfitlovID:1532}, etc{})

    # for _, v in enumerate(userInfosList):
    #     listID = v['crossfitlovID']

    for i, v in enumerate(userInfosList):
        usr_lambda = v["crossfitlovID"]
        print(v["email"])
        if str(usr_lambda) == mec_courant:
            userInfosList.remove(v)
        elif userInfosList in idBDDSwipeList:##idBDDSwipe.contains(v["crossfitlovID"]:
            userInfosList.remove(v)
        else:
            continue
    # for i, v in enumerate(idBDDSwipe):
    #     for u, t in enumerate(listID):

    #         if v == t:
    #             listID.remove(u)


    return jsonify(userInfosList)

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