import { Component, OnInit } from '@angular/core';
import {NgbModal, ModalDismissReasons} from '@ng-bootstrap/ng-bootstrap';
import { AuthService } from 'src/app/services/auth.services';





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

  constructor(private modalService: NgbModal, private authService: AuthService) {

   }

  ngOnInit(): void {
  }

  signin(): void {
    this.authService.login(this.login, this.password).subscribe(
      () => {
        console.log('Enregistrement terminé !');
        /*const user: User = {firstname: "yeahh", age: 22, crossfitlovID : 1, };
        sessionStorage.setItem('token', JSON.stringify('response: auth'));
        this.subject.next(user);
        this.router.navigate(["/home"]);*/
      },
      (error) => {
        console.log('Erreur ! : ' + error);
      }
    );
  }

  isLoginValid(): boolean{
      return (!this.login || !this.password);
  }

} 
