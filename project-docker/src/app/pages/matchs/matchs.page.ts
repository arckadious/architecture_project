import { Component, OnInit } from '@angular/core';
import { Person } from '../../services/persons.service';
import { MessagesService  } from '../../services/messages.service';

@Component({
  selector: 'app-matchs',
  templateUrl: './matchs.page.html',
  styleUrls: ['./matchs.page.scss'],
})
export class MatchsPage implements OnInit {

  persons : any[]; 
  constructor(private Person: Person,private messageService: MessagesService ) { }

  ngOnInit() {

    this.persons = this.Person.matchs;

  }

  goToRoom(id:number) : void{

    this.messageService.currentRoom = id;
    this.messageService.getMessages();
    
  }
  

}
