import { Component, OnInit } from '@angular/core';
import { ModalController } from '@ionic/angular';
import { MatchModalComponent} from '../match-modal/match-modal.component'
import { Person } from '../../services/persons.service';
import { MatchsService } from '../../services/matchs.service';



@Component({
  selector: 'app-swipe',
  templateUrl: './swipe.page.html',
  styleUrls: ['./swipe.page.scss'],

  
})
export class SwipePage implements OnInit {

  currentIndex: number;
  i: number ;
  results = [];
  persons : any[]; 

  // persons = [
  //   {
  //     name: 'Alexis Ren',
  //     age: 24,
  //     image: '../../assets/alexis-ren.jpg',
  //     hasSwiped: false,
  //     visible: true
  //   },
  //   {
  //     name: 'Megan Fox',
  //     age: 34,
  //     image: '../../assets/megan-fox.jpg',
  //     hasSwiped: false,
  //     visible: true
  //   },
  //   {
  //     name: 'Scarlett Johansson',
  //     age: 36,
  //     image: '../../assets/scarlett-jo.jpg',
  //     hasSwiped: false,
  //     visible: true
  //   },
  //   {
  //     name: 'Beyonce',
  //     age: 38,
  //     image: '../../assets/beyonce.jpg',
  //     hasSwiped: true,
  //     visible: true
  //   },
  // ];

  constructor(public modalController: ModalController,private Person: Person,private matchService: MatchsService ) {
    // this.currentIndex = this.persons.length - 1;
  }
  ngOnInit(): void {
    this.persons = this.Person.persons;
    console.log(this.persons)
    this.currentIndex = this.persons.length - 1;
    console.log(this.currentIndex);


  }

  async showModal(){
    const modal = await this.modalController.create({
      component: MatchModalComponent,
      componentProps: {
        data: this.persons[this.currentIndex].name
      }
    })
    await modal.present()
  }
  swiped(event: any, index: number) {
    this.persons[index].visible = false;
    this.results.push(this.persons[index].name + ' swiped ' + event);
    this.currentIndex--;
    this.i++;
  }


  swipeleft() {
    this.persons[this.currentIndex].visible = false;
    this.results.push(this.persons[this.currentIndex].name + ' swiped false');
    this.currentIndex--;
    this.i++;

  }

  swiperight() {
    this.persons[this.currentIndex].visible = false;
    this.results.push(this.persons[this.currentIndex].name + ' swiped true');
    if(this.persons[this.currentIndex].hasSwiped == true){
      this.Person.matchs.push(this.persons[this.currentIndex]);
      console.log(this.Person.matchs)
      this.matchService.swipeId = this.persons[this.currentIndex].id;
      this.matchService.userId = 1
      this.matchService.sendMatch()
      this.showModal()
        }
    this.currentIndex--;
    this.i++;

  }

}

