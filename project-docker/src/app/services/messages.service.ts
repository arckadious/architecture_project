import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Messages } from '../models/messages';


@Injectable({
  providedIn: 'root'
})
export class MessagesService {
  currentRoom :  number;
  messages : Messages[]
  constructor(private http: HttpClient) { 

   
  }
  
  getMessages(){
    let headers = {
      'Content-Type': 'application/json',
      'X-Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6IjEiLCJleHAiOjE2MTQ1NDc3MTB9.XwmNI3kXRxwBg6La_LibyoYDao7jR3NdDbMQpENPV4I',
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
      }
    );  }
}
