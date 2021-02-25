import { Component, OnInit } from '@angular/core';
import { Person } from '../services/persons.service';
import { HttpClient, HttpParams } from '@angular/common/http';


@Component({
  selector: 'app-registration',
  templateUrl: './registration.page.html',
  styleUrls: ['./registration.page.scss'],
})
export class RegistrationPage implements OnInit {

person:Person = new Person();
  


  constructor(private http: HttpClient) { }

  ngOnInit() {
  }

  userRegister(){
    let data = {
      "login": this.person.login,
      "password": this.person.password,
      "firstname": this.person.firstname,
      "boxCity": this.person.boxCity,
      "biography": this.person.biography,
      "job": this.person.job,
      "gender": this.person.gender,
      "email": this.person.email,
    }
    
    JSON.stringify(data) 

    this.http.post('https://httpclient-demo.firebaseio.com/appareils.json', data)
        .subscribe(
          () => {
            console.log('Enregistrement terminé !');
          },
          (error) => {
            console.log('Erreur ! : ' + error);
          }
        );
  }

}
