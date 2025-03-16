import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Post } from 'src/app/Models/post';

@Component({
  selector: 'app-view-all-post',
  templateUrl: './view-all-post.component.html',
  styleUrls: ['./view-all-post.component.scss']
})
export class ViewAllPostComponent implements OnInit {

  
  @Input() posts: any;
 
  
  constructor(private router: Router) {console.log(this.posts) }

  ngOnInit(): void {
    console.log(this.posts)
  }

  goToPost(id: number): void {
    console.log("read")
    this.router.navigateByUrl('/view-post/' + id);
    console.log(id)
  }

}
