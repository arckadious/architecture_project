import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class Person {

  matchs = [];  
  sexe = ["Boy","Girl"];
  login:string;
  firstname:string;
  email:string;
  password:string;
  boxCity:string;
  biography:string;
  job:string;
  gender:string;
  createAt:string;

  persons = [
    {
      id: 1,
      name: 'Alexis',
      status: 'En ligne',
      img :"../../assets/alexis-ren.jpg",
      age: 24,
      hasSwiped: false,
      visible: true
      
    },
    {
      id: 2,
      name: 'Megan',
      status: 'Hors-ligne',
      img :'../../assets/megan-fox.jpg',
      age: 34,
      hasSwiped: false,
      visible: true

    },
    {
      id: 3,
      name: 'Scarlett',
      status: 'En ligne',
      img :'../../assets/scarlett-jo.jpg',
      age: 34,
      hasSwiped: true,
      visible: true

    }
  ];
  constructor() { }
}
