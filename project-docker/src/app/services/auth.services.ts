import { Injectable } from '@angular/core'
import { User } from '../models/user.model';
import { Observable, BehaviorSubject} from 'rxjs';
import { Router } from '@angular/router';
import { HttpClient, HttpParams } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { SigninData } from '../models/signinData.model';


@Injectable({
    providedIn: 'root'
})

export class AuthService {
    private subject: BehaviorSubject<User> = new BehaviorSubject<User>(null); //subject

    get user(): Observable<User> {   
        return this.subject.asObservable();
    }

    constructor(private router: Router, private http: HttpClient) {
        //subject.next execute l'observable
        const u = JSON.parse(sessionStorage.getItem('USER'));
        if(u != null) {
            // if(/*si le token n'est pas périmé*/) {
            //     this.subject.next(u);
            // }
        }

    }

    login(login: string, password: string) : Observable<SigninData> {
        let data = {
          "login": login,
          "password": password
        }
        
        //appel de la brique auth-api pour s'authentifier
        let headers = {
            'Content-Type': 'application/json',
            'Authorization': 'Basic ' + btoa(environment.auth_api_config.auth_api_basicauth_login + ':' + environment.auth_api_config.auth_api_basicauth_password)
        }
          
      
        return this.http.post<SigninData>(environment.auth_api_config.auth_api_URL+"/signin", JSON.stringify(data), { headers });
      
               
    }


    logout(): void {
        sessionStorage.removeItem('token');
        this.subject.next(null);
        this.router.navigate(["/home"]);
    }
}