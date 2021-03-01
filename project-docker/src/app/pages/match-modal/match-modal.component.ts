import { Component, OnInit,Input } from '@angular/core';
import { ModalController } from '@ionic/angular';
import { MessagesService  } from '../../services/messages.service';

@Component({
  selector: 'app-match-modal',
  templateUrl: './match-modal.component.html',
  styleUrls: ['./match-modal.component.scss'],
})
export class MatchModalComponent implements OnInit {
  @Input() data: any;
  firstMessage:String;
  constructor(public modalController: ModalController,private messageService: MessagesService) { }

  ngOnInit() {}

  goToRoom(id:number) : void{

    this.messageService.currentRoom = id;
  }

}

  


