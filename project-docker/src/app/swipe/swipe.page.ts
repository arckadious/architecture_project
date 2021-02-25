import { Component, OnInit } from '@angular/core';
import { ModalController } from '@ionic/angular';

@Component({
  selector: 'app-swipe',
  templateUrl: './swipe.page.html',
  styleUrls: ['./swipe.page.scss'],

  
})
export class SwipePage implements OnInit {

  currentIndex: number;
  results = [];
  avatars = [
    {
      name: 'Alexis Ren',
      age: 24,
      image: '../../assets/alexis-ren.jpg',
      hasSwiped: true,
      visible: true
    },
    {
      name: 'Megan Fox',
      age: 34,
      image: '../../assets/megan-fox.jpg',
      hasSwiped: false,
      visible: true
    },
    {
      name: 'Scarlett Johansson',
      age: 36,
      image: '../../assets/scarlett-jo.jpg',
      hasSwiped: false,
      visible: true
    },
    {
      name: 'Beyonce',
      age: 38,
      image: '../../assets/beyonce.jpg',
      hasSwiped: false,
      visible: true
    },
  ];

  constructor(public modalController: ModalController) {
    this.currentIndex = this.avatars.length - 1;
    console.log(this.currentIndex);
  }
  ngOnInit(): void {
  }

  swiped(event: any, index: number) {
    this.avatars[index].visible = false;
    this.results.push(this.avatars[index].name + ' swiped ' + event);
    this.currentIndex--;
  }


  swipeleft() {
    this.avatars[this.currentIndex].visible = false;
    this.results.push(this.avatars[this.currentIndex].name + ' swiped false');
    this.currentIndex--;
  }

  swiperight() {
    this.avatars[this.currentIndex].visible = false;
    this.results.push(this.avatars[this.currentIndex].name + ' swiped true');
    this.currentIndex--;
  }

}
