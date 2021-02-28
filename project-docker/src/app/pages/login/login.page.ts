import { Component, OnInit } from '@angular/core';
import {NgbModal, ModalDismissReasons} from '@ng-bootstrap/ng-bootstrap';
import { AuthService } from 'src/app/services/auth.services';
import { Router } from '@angular/router';






@Component({
  selector: 'app-login',
  templateUrl: './login.page.html',
  styleUrls: ['./login.page.scss'],
})
export class LoginPage implements OnInit {

  //person : Person = new Person();
  closeResult = '';
  login : string = '';
  password:string = '';

  constructor(private modalService: NgbModal, private authService: AuthService,private router: Router) {

   }

  ngOnInit(): void {
  }

  signin(): void {
    this.authService.login(this.login, this.password);
    

  }

  isLoginValid(): boolean{
      return (!this.login || !this.password);
  }

} 
