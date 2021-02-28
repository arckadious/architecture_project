import { Component, OnInit } from '@angular/core';
import { Messages } from '../../models/messages';
import { MessagesService  } from '../../services/messages.service';


@Component({
  selector: 'app-chat',
  templateUrl: './chat.page.html',
  styleUrls: ['./chat.page.scss'],
})
export class ChatPage implements OnInit {

  messages : Messages[]

  constructor(private messageService: MessagesService) { }

  ngOnInit() {

    this.messages = this.messageService.messages;
  }

}
