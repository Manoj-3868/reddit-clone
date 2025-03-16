import { Component, HostListener, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Post } from 'src/app/Models/post';
import { PostService } from 'src/app/Services/post.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {
  posts: Array<Post> = [];

  constructor(private postService: PostService,router :Router) { 
   let b= localStorage.getItem("username")+''
    this.postService.getAllPosts(b).subscribe(post => {
      this.posts = post;
    });
  }

  
  ngOnInit(): void {
    
  }
 
}
