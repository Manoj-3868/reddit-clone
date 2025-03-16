import { Component, Input, OnInit } from '@angular/core';
import { Post } from 'src/app/Models/post';
import { Vote } from 'src/app/Models/vote';
import { AuthService } from 'src/app/Services/auth.service';
import { PostService } from 'src/app/Services/post.service';
import { VoteService } from 'src/app/Services/vote.service';

@Component({
  selector: 'app-manage-vote',
  templateUrl: './manage-vote.component.html',
  styleUrls: ['./manage-vote.component.scss']
})
export class ManageVoteComponent implements OnInit {

  @Input() post: any;

  votePayload: Vote;
  b :string;
  upvoteColor: string | undefined;
  downvoteColor: string | undefined;
  isLoggedIn: boolean | undefined;

  constructor(private voteService: VoteService,private postService: PostService) {
    this.b = localStorage.getItem("loggedin")+''
    this.votePayload = {
      votetype:0,
      postid: this.post?._id,
      username: localStorage.getItem("username")+''
    }  
  }

  ngOnInit(): void {
    
  }

  upvotePost() {
    this.votePayload.votetype = 1;
    if (this.b=="true"){
      this.vote();
    }else{
      alert("login to vote")
    }

    this.downvoteColor = '';
  }

  downvotePost() {
    this.votePayload.votetype = -1;
    if (this.b=="true"){
      this.vote();
    }else{
      alert("login to vote")
    }
    this.upvoteColor = '';
  }

  private vote() {
    this.votePayload.postid = this.post?._id;
    this.voteService.vote(this.votePayload).subscribe(() => {
      this.updateVoteDetails();
    });
  }

  private updateVoteDetails() {
    this.postService.getPost(this.post._id).subscribe(post => {
      this.post = post;      
    });
  }

}
