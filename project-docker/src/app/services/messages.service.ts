import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Messages } from '../models/messages';
import { User } from '../models/user.model';
import { AuthService } from 'src/app/services/auth.services';

@Injectable({
  providedIn: 'root'
})
export class MessagesService {
  currentRoom :  number;
  userId :  number;
  swipeId :  number;



  messages : Messages[]
  user: User
  constructor(private http: HttpClient, private auth: AuthService) { 
    this.auth.user.subscribe(data => this.user = data)

   
  }
  
  getMessages(){
    let headers = {
      'Content-Type': 'application/json',
      'X-Authorization': 'Bearer '+this.user.tokenInfos.value,
      'Authorization': 'Basic ' + btoa(environment.message_api_config.basicauth_login + ':' + environment.message_api_config.basicauth_password)
    }

    let data = [{
      room_id :  this.currentRoom.toString()
    }]
    this.http.post<Messages[]>(environment.message_api_config.URL+"/api/message", JSON.stringify(data), { headers }).subscribe(
      (msg) => {
        this.messages = msg;
        console.log(msg);
      },
      (error) => {
        console.log('Erreur ! : ' + error);
        alert("Vous avez été déconnecté, veuillez vous reconnecter.")
        this.auth.logout();
      }
    ); 
   }

    sendMatch(){

      let data = [{
        usr_id :  this.userId.toString(),
        swipe_id :this.swipeId.toString() 
      }]
      let headers = {
        'Content-Type': 'application/json',
        'X-Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6IjEiLCJleHAiOjE2MTQ1NDc3MTB9.XwmNI3kXRxwBg6La_LibyoYDao7jR3NdDbMQpENPV4I',
        'Authorization': 'Basic ' + btoa(environment.match_api_config.basicauth_login + ':' + environment.match_api_config.basicauth_password)
      }
      this.http.post<Messages[]>(environment.match_api_config.URL+"/api/match", JSON.stringify(data), { headers }).subscribe(
        (msg) => {
          this.messages = msg;
          console.log(msg);
        },
        (error) => {
          console.log('Erreur ! : ' + error);
          alert("Vous avez été déconnecté, veuillez vous reconnecter.")
          this.auth.logout();
        }
      ); 


    }
}
