import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { AuthService } from 'src/app/services/auth.services';
import { UserInfos } from 'src/app/models/user.model';
import {formatDate} from '@angular/common';


@Component({
  selector: 'app-registration',
  templateUrl: './registration.page.html',
  styleUrls: ['./registration.page.scss'],
})
export class RegistrationPage implements OnInit {  

  login : string
  password : string
  userInfos : UserInfos

  constructor(private http: HttpClient, private authService: AuthService) { 
    this.userInfos = {} as UserInfos;
  }

  ngOnInit() {
  }

  userRegister(){
    let time = new Date()
    let hours = time.getHours() < 10 ? '0'+time.getHours().toString() : time.getHours().toString();
    let minutes = time.getMinutes() < 10 ? '0'+time.getMinutes().toString() : time.getMinutes().toString();
    let seconds = time.getSeconds() < 10 ? '0'+time.getSeconds().toString() : time.getSeconds().toString();
 
    this.userInfos.createdAt = formatDate(time, 'yyyy-MM-dd', 'en')+' '+hours+':'+minutes+':'+seconds;
    console.log(this.userInfos.gender)
    this.authService.signup(this.userInfos, this.login, this.password);
  }
  
}
