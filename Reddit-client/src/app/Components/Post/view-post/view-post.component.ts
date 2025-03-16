import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Post } from 'src/app/Models/post';
import { Comment } from 'src/app/Models/comment';
import { PostService } from 'src/app/Services/post.service';
import { ActivatedRoute, Router } from '@angular/router';
import { CommentService } from 'src/app/Services/comment.service';

@Component({
  selector: 'app-view-post',
  templateUrl: './view-post.component.html',
  styleUrls: ['./view-post.component.scss']
})
export class ViewPostComponent implements OnInit {

  postId: number;
  post: Post ={
    _id: 0,
    postname: '',
    url: '',
    description: '',
    votecount: 0,
    username: '',
    subredditname: '',
    commentcount: 0,
    upvote: false,
    downvote: false
  };
  commentForm: FormGroup;
  commentPayload: Comment;
  comments: Comment[]=[];

  constructor(private postService: PostService, private activateRoute: ActivatedRoute,
    private commentService: CommentService, private router: Router) {
    this.postId = this.activateRoute.snapshot.params['id'];

    this.commentForm = new FormGroup({
      text: new FormControl('', Validators.required)
    });
    this.commentPayload = {
      text: '',
      postid: parseInt(this.postId+'')
    };
  }

  ngOnInit(): void {
    this.getPostById();
    this.getCommentsForPost();
  }

  postComment() {
    console.log(this.commentPayload)
    this.commentPayload.username = localStorage.getItem("username")+''
    this.commentPayload.text = this.commentForm.get('text')?.value;
    this.commentService.postComment(this.commentPayload).subscribe(data => {
      this.commentForm.get('text')?.setValue('')
      this.getCommentsForPost();
    })
  }

  private getPostById() {

    this.postService.getPost(this.postId).subscribe(data => {
      this.post = data;
    });
  }

  private getCommentsForPost() {
    this.commentService.getAllCommentsForPost(this.postId).subscribe(data => {
      this.comments = data;
    });
  }

}
