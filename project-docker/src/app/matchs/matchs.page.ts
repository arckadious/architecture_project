import { Component, OnInit } from '@angular/core';
import { Person } from '../services/persons.service';

@Component({
  selector: 'app-matchs',
  templateUrl: './matchs.page.html',
  styleUrls: ['./matchs.page.scss'],
})
export class MatchsPage implements OnInit {

  persons : any[]; 
  constructor(private Person: Person) { }

  ngOnInit() {

    this.persons = this.Person.persons;

  }

}
