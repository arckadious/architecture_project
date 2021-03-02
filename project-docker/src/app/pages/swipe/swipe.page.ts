import { Component, OnInit } from '@angular/core';
import { ModalController } from '@ionic/angular';
import { MatchModalComponent} from '../match-modal/match-modal.component'
//import { Person } from '../../services/persons.service';
import { MatchsService } from '../../services/matchs.service';
import { NumberValueAccessor } from '@angular/forms';

import { User, UserInfos } from '../../models/user.model';
import { IsMatch } from '../../models/matchs';
import { AuthService } from 'src/app/services/auth.services';


@Component({
  selector: 'app-swipe',
  templateUrl: './swipe.page.html',
  styleUrls: ['./swipe.page.scss'],

  
})
export class SwipePage implements OnInit {

  currentIndex: number;
  i: number ;
  results = [];
  swipeNumber : number = 10;
  isSub:boolean
  user: User

  persons : UserInfos[]

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

  constructor(public modalController: ModalController, private authService: AuthService, private matchService: MatchsService ) {
    // this.currentIndex = this.persons.length - 1;
    authService.user.subscribe(data => this.user = data);
  }
  ngOnInit(): void {
    this.matchService.getUsers().subscribe(
      (element) => {
        const userList = element;
        this.persons = userList;
        
        console.log(this.persons)
        this.currentIndex = this.persons.length -1;
        console.log(this.currentIndex);
        console.log(this.matchService.isSub)
      }
    );


    this.matchService.isSubscribed();
    
  }

  async showModal(){
    const modal = await this.modalController.create({
      component: MatchModalComponent,
      componentProps: {
        data: this.persons[this.currentIndex].firstname
      }
    })
    await modal.present()
  }
  swiped(event: any, index: number) {
    if(this.swipeNumber == 0 || this.currentIndex == 0){
      return;
    }
    this.results.push(this.persons[index].firstname + ' swiped ' + event);
    this.currentIndex--;
    if(this.matchService.isSub == false){

      this.swipeNumber-=5
    }else{

      this.swipeNumber--
    }    this.i++;
  }


  swipeleft() {
    if(this.swipeNumber == 0 || this.currentIndex == -1){
      return;
    }
    this.results.push(this.persons[this.currentIndex].firstname + ' swiped false');
    this.currentIndex--;
    if(this.matchService.isSub == false){

      this.swipeNumber-=5
    }else{

      this.swipeNumber--
    }   
    console.log(this.swipeNumber)

    this.i++;


  }

  swiperight() {
    if(this.swipeNumber == 0 || this.currentIndex == -1){
      return;
    }

    this.matchService.sendSwipe(this.user.userInfos.crossfitlovID, this.persons[this.currentIndex].crossfitlovID).subscribe(
      (element) => {
        const ismatch = element;
        console.log(ismatch.ismatch);


        this.results.push(this.persons[this.currentIndex].firstname + ' swiped true');
        if(ismatch.ismatch == "Yes"){
          this.showModal()
        }
        if(this.matchService.isSub == false){
    
          this.swipeNumber-=5
        }else{
    
          this.swipeNumber--
        }
        console.log(this.swipeNumber)
        this.currentIndex--;
        this.i++;



      },
      (error) => {
        console.log('Erreur ! : ' + error);
        alert("Vous avez été déconnecté, veuillez vous reconnecter.")
        //this.auth.logout();
      }
    )

    
 
  }

  sabonner(){

    this.matchService.userSubscribe();
  }


 

}

