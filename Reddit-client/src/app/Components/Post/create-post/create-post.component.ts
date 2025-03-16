import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { Postrequest } from 'src/app/Models/postrequest';
import { Subreddit } from 'src/app/Models/subreddit';
import { PostService } from 'src/app/Services/post.service';
import { SubredditService } from 'src/app/Services/subreddit.service';

@Component({
  selector: 'app-create-post',
  templateUrl: './create-post.component.html',
  styleUrls: ['./create-post.component.scss']
})
export class CreatePostComponent implements OnInit {

  createPostForm: FormGroup;
  postPayload: Postrequest;
  subreddits: Array<Subreddit>=[];

  constructor(private router: Router, private postService: PostService,
    private subredditService: SubredditService) {

    this.postPayload = {
      postname: '',
      url: '',
      description: '',
      subredditname: '',
      username:''
    }
    this.createPostForm = new FormGroup({
      postName: new FormControl('', Validators.required),
      subredditName: new FormControl('', Validators.required),
      url: new FormControl('', Validators.required),
      description: new FormControl('', Validators.required),
    });
  }

  ngOnInit() {
    
    this.subredditService.getAllSubreddits().subscribe((data) => {
      this.subreddits = data;
    })
  }

  createPost() {
    this.postPayload.postname = this.createPostForm.get('postName')?.value;
    this.postPayload.subredditname = this.createPostForm.get('subredditName')?.value;
    this.postPayload.url = this.createPostForm.get('url')?.value;
    this.postPayload.description = this.createPostForm.get('description')?.value;
    let b = localStorage.getItem("username")
    this.postPayload.username = b+""
    this.postService.createPost(this.postPayload).subscribe((data) => {
      this.router.navigateByUrl('/');
    })
  }

  discardPost() {
    this.router.navigateByUrl('/');
  }

}
