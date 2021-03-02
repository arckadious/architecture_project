import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Messages } from '../models/messages';
import { Matchs } from '../models/matchs';
import { IsMatch } from '../models/matchs';

import { User } from '../models/user.model';
import { AuthService } from 'src/app/services/auth.services';
import { Abonnement } from '../models/abonnement';
import { UserInfos } from '../models/user.model';
import { Observable, BehaviorSubject } from 'rxjs';


@Injectable({
  providedIn: 'root'
})


export class MatchsService {

  match:Matchs;
  isSub: boolean;
  sub: Abonnement;

  private user: User;



  constructor(private auth: AuthService, private http: HttpClient) {
      this.auth.user.subscribe(data => this.user = data)
  }
  sendSwipe(userId, swipeId) : Observable<IsMatch>{

    let data = [{
      "usr_id" :  userId.toString(),
      "swipe_with" :swipeId.toString()
    }]
    let headers = {
      'Content-Type': 'application/json', 
      'X-Authorization': 'Bearer '+this.user.tokenInfos.value,
      'Authorization': 'Basic ' + btoa(environment.match_api_config.basicauth_login + ':' + environment.match_api_config.basicauth_password)
    }
    
    return this.http.put<IsMatch>(environment.match_api_config.URL+"/api/matches", JSON.stringify(data), { headers }); 


  }


  userSubscribe(){

    let data = [{
      "id_usr" :  this.user.userInfos.crossfitlovID.toString(),
      "y_or_n" :  1
    }]
    let headers = {
      'Content-Type': 'application/json', 
      'X-Authorization': 'Bearer '+this.user.tokenInfos.value,
      'Authorization': 'Basic ' + btoa(environment.abonnement_api_config.basicauth_login + ':' + environment.abonnement_api_config.basicauth_password)
    }

    
    this.http.put<Abonnement>(environment.abonnement_api_config.URL+"/api/abonnement", JSON.stringify(data), { headers }).subscribe(
      (sub) => {
        this.sub = sub;
        console.log(this.sub.Abonnement);

      }
      
    ); 


  }

  isSubscribed(){

    let data = [{
      "id_usr" :  this.user.userInfos.crossfitlovID.toString()
      
    }]
    let headers = {
      'Content-Type': 'application/json', 
      'X-Authorization': 'Bearer '+this.user.tokenInfos.value,
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
      }
      
    ); 


  }

  getUsers() : Observable<UserInfos[]>{


    let data = {
      "usr_id" :  this.user.userInfos.crossfitlovID.toString(),
    }
    let headers = {
      'Content-Type': 'application/json', 
      'X-Authorization': 'Bearer '+this.user.tokenInfos.value,
      'Authorization': 'Basic ' + btoa(environment.match_api_config.basicauth_login + ':' + environment.match_api_config.basicauth_password)
    }
    
    return this.http.post<UserInfos[]>(environment.match_api_config.URL+"/api/show", JSON.stringify(data), { headers }); 
  }

}
