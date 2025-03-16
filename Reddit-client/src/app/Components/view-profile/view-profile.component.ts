import { Component, OnInit } from '@angular/core';
import { Post } from 'src/app/Models/post';
import { Comment } from 'src/app/Models/comment';
import { PostService } from 'src/app/Services/post.service';
import { CommentService } from 'src/app/Services/comment.service';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-view-profile',
  templateUrl: './view-profile.component.html',
  styleUrls: ['./view-profile.component.scss']
})
export class ViewProfileComponent implements OnInit {

  name: string;
  posts: Post[] =[];
  comments: Comment[]=[];
  postLength: number = 0;
  commentLength: number =0;

  constructor(private activatedRoute: ActivatedRoute, private postService: PostService,
    private commentService: CommentService) {
    this.name = this.activatedRoute.snapshot.params['name'];
      console.log(this.name)
    this.postService.getAllPostsByUser(this.name).subscribe(data => {
      this.posts = data;
      this.postLength = data.length;
    });
    this.commentService.getAllCommentsByUser(this.name).subscribe(data => {
      this.comments = data;
      this.commentLength = data.length;
    });
  }

  ngOnInit(): void {
  }


}
