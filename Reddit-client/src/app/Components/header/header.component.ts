import { Component, HostListener, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {
  isLoggedIn:Boolean = false
  username:string|null=''
  @HostListener('window:keydown',['$event']) spacevent(event:any){
    if (event.key == "A" ){
      this.router.navigate(['/create-post'])
    }
    if (event.key == "S" ){
      this.router.navigate(['/create-subreddit'])
    }

  }
  constructor(private router:Router) {
   let a = localStorage.getItem("loggedin")
   let b = localStorage.getItem("username")
   if(a == "true"){
      this.isLoggedIn =true
      this.username=b
   }
   }

  ngOnInit(): void {
  }

  logout(){
    localStorage.clear()
    console.log("loggout")
    
    this.router.navigateByUrl('/login');
  }
  goToUserProfile() {
    this.router.navigateByUrl('/user-profile/' + this.username);
  }
}
