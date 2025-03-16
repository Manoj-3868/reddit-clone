import { Component, OnInit } from '@angular/core';
import { Subreddit } from 'src/app/Models/subreddit';
import { SubredditService } from 'src/app/Services/subreddit.service';


@Component({
  selector: 'app-viewall-subreddit',
  templateUrl: './viewall-subreddit.component.html',
  styleUrls: ['./viewall-subreddit.component.scss']
})
export class ViewallSubredditComponent implements OnInit {
  
  subreddits: Array<Subreddit>=[];
  constructor(private subredditService: SubredditService) { 
    this.subredditService.getAllSubreddits().subscribe(data => {
      this.subreddits = data;
    })
  }

  ngOnInit() {
    }
}
