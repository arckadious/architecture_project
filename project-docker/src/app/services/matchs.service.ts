import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Messages } from '../models/messages';
import { Matchs } from '../models/matchs';
import { User } from '../models/user.model';
import { AuthService } from 'src/app/services/auth.services';
import { Abonnement } from '../models/abonnement';



@Injectable({
  providedIn: 'root'
})


export class MatchsService {
  userId :  number;
  swipeId :  number;
  match:Matchs;
  isSub: boolean;
  sub: Abonnement;

  private user: User;

  constructor(private auth: AuthService, private http: HttpClient) {
      this.auth.user.subscribe(data => this.user = data)
  }
  sendMatch(){

    let data = [{
      "usr_id" :  this.userId.toString(),
      "swipe_with" :this.swipeId.toString()
    }]
    let headers = {
      'Content-Type': 'application/json', 
      'X-Authorization': 'Bearer '+localStorage.getItem("token-CL").toString(),
      'Authorization': 'Basic ' + btoa(environment.match_api_config.basicauth_login + ':' + environment.match_api_config.basicauth_password)
    }
    
    this.http.put<Matchs>(environment.match_api_config.URL+"/api/matches", JSON.stringify(data), { headers }).subscribe(
      (match) => {
        this.match= match;
        console.log(match);
      },
      (error) => {
        console.log('Erreur ! : ' + error);
        alert("Vous avez été déconnecté, veuillez vous reconnecter.")
        //this.auth.logout();
      }
    ); 


  }

  userSubscribe(){

    let data = [{
      "usr_id" :  this.user.userInfos.crossfitlovID.toString(),
      "swipe_with" : "y"
    }]
    let headers = {
      'Content-Type': 'application/json', 
      'X-Authorization': 'Bearer '+localStorage.getItem("token-CL").toString(),
      'Authorization': 'Basic ' + btoa(environment.abonnement_api_config.basicauth_login + ':' + environment.abonnement_api_config.basicauth_password)
    }
    
    this.http.put<Abonnement>(environment.abonnement_api_config.URL+"/api/matches", JSON.stringify(data), { headers }).subscribe(
      (sub) => {
        console.log(sub)

      },
      (error) => {
        console.log('Erreur ! : ' + error);
        alert("Vous avez été déconnecté, veuillez vous reconnecter.")
        //this.auth.logout();
      }
    ); 


  }

  isSubscribed(){

    let data = [{
      "id_usr" :  this.user.userInfos.crossfitlovID.toString()
      
    }]
    let headers = {
      'Content-Type': 'application/json', 
      'X-Authorization': 'Bearer '+localStorage.getItem("token-CL").toString(),
      'Authorization': 'Basic ' + btoa(environment.abonnement_api_config.basicauth_login + ':' + environment.abonnement_api_config.basicauth_password)
    }
    
    this.http.post<Abonnement>(environment.abonnement_api_config.URL+"/api/abonnement", JSON.stringify(data), { headers }).subscribe(
      (sub) => {
        this.sub = sub;
        console.log(this.sub.Abonnement);
        if(this.sub.Abonnement == "y"){

          this.isSub = true;
          console.log(this.isSub)

        }else{

          this.isSub = false;
          console.log(this.isSub)


        }
      },
      (error) => {
        console.log('Erreur ! : ' + error);
        alert("Vous avez été déconnecté, veuillez vous reconnecter.")
        //this.auth.logout();
      }
    ); 


  }
}
