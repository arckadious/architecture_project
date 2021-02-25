import { Component, OnInit } from '@angular/core';
import {NgbModal, ModalDismissReasons} from '@ng-bootstrap/ng-bootstrap';
import { Person } from '../services/persons.service';
import { HttpClient, HttpParams } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';




@Component({
  selector: 'app-login',
  templateUrl: './login.page.html',
  styleUrls: ['./login.page.scss'],
})
export class LoginPage implements OnInit {

  person:Person = new Person();



  closeResult = '';
  login : string = '';
  password:string = '';

  constructor(private modalService: NgbModal,private http: HttpClient) {

   }

  ngOnInit(): void {
  }

  open(content: any) {
    this.modalService.open(content, {ariaLabelledBy: 'modal-basic-title'}).result.then((result) => {
      this.closeResult = `Closed with: ${result}`;
    }, (reason) => {
      this.closeResult = `Dismissed ${this.getDismissReason(reason)}`;
    });
  }
  
  private getDismissReason(reason: any): string {
    if (reason === ModalDismissReasons.ESC) {
      return 'by pressing ESC';
    } else if (reason === ModalDismissReasons.BACKDROP_CLICK) {
      return 'by clicking on a backdrop';
    } else {
      return  `with: ${reason}`;
    }
  }

  isLoginValid(): boolean{
      return (!this.person.login || !this.person.password);
  }

 
  // addUser(data): Observable<Person> {
  //   return this.http.post<Person>('http://www.example.com/', JSON.stringify(data), )
  //   .pipe(
  //     retry(1),
  //     catchError(this.processError)
  //   )
  // } 
  userLogin(){
    let data = {
      "login": this.person.login,
      "password": this.person.password


    }
    
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
