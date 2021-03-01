import { Injectable } from '@angular/core'
import { Observable, BehaviorSubject} from 'rxjs';
import { Router } from '@angular/router';
import { HttpClient, HttpParams } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { User } from '../models/user.model';
import { UserInfos } from '../models/user.model';


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
            this.subject.next(u);
        }

    }

    login(login: string, password: string) : void {
        let data = {
          "login": login,
          "password": password,
          "getInfos": true
        }
        
        //appel de la brique auth-api pour s'authentifier
        let headers = {
            'Content-Type': 'application/json',
            'Authorization': 'Basic ' + btoa(environment.auth_api_config.basicauth_login + ':' + environment.auth_api_config.basicauth_password)
        }
          
      
        this.http.post<User>(environment.auth_api_config.URL+"/signin", JSON.stringify(data), { headers }).subscribe(
            (element) => {
              console.log('Enregistrement terminé !');
              const user: User = element;
              localStorage.setItem("token-CL", user.tokenInfos.value);
              user.tokenInfos.value = "******"
              this.subject.next(user);
              this.router.navigate(['/swipe']);
            },
            (error) => {
              alert("Utilisateur ou mot de passe incorrect.")
              console.log('Erreur ! : ' + error);
            }
          );           
    }

    signup(userInfos: UserInfos, login: string, password: string) : void {
        
        let data = {
            credentialsData: {
                login: login,
                password : password,
            },
            userInfos: userInfos      
        }
        //appel de la brique auth-api pour s'authentifier
        let headers = {
            'Content-Type': 'application/json',
            'Authorization': 'Basic ' + btoa(environment.auth_api_config.basicauth_login + ':' + environment.auth_api_config.basicauth_password)
        }
          
      
        this.http.put<User>(environment.auth_api_config.URL+"/signup", JSON.stringify(data), { headers }).subscribe(
            () => {
              console.log('inscription terminé !');
              alert("Inscription effectué, vous pouvez maintenant vous connecter.")

              this.router.navigate(['/login']);
            },
            (error) => {
              alert("Certaines informations semblent incorrectes, Veuillez remplir les champs correctement")
              console.log('Erreur ! : ' + error);

            }
          );           
    }


    logout(): void {
        localStorage.removeItem('token-CL');
        this.subject.next(null);
        this.router.navigate(["/home"]);
    }
}