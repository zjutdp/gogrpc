import { Component, OnInit } from '@angular/core';
import { DataService } from '../data.service'
import { Observable } from 'rxjs'
import { ActivatedRoute } from '@angular/router'
import { User } from '../model/user'

@Component({
  selector: 'app-details',
  templateUrl: './details.component.html',
  styleUrls: ['./details.component.scss']
})
export class DetailsComponent implements OnInit {

  userId: string;
  user$: Object;

  constructor(private data: DataService, private route: ActivatedRoute) { 
    this.route.params.subscribe(params => this.userId = params.id)
  }

  ngOnInit() {
    this.user$ = JSON.parse(sessionStorage.getItem(this.userId))

    if (this.user$ == null) {
      this.data.getUser(this.userId).subscribe(
          data => {
            this.user$ = new User(data.id, data.name, data.email, data.phone)
            sessionStorage.setItem(this.userId, JSON.stringify(this.user$))
          })
    }

    console.log('inited')
  }

}
