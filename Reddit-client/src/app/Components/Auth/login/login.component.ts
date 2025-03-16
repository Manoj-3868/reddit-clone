import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { Login } from 'src/app/Models/login';
import { AuthService } from 'src/app/Services/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  loginForm: FormGroup;
  loginRequestPayload: Login;
  // registerSuccessMessage: string;
  // isError: boolean;

  constructor(private router: Router,private authService:AuthService) {
    this.loginRequestPayload = {
      username: '',
      password: ''
    };
    this.loginForm = new FormGroup({
      username: new FormControl('', Validators.required),
      password: new FormControl('', Validators.required)
    });
    let a = localStorage.getItem("loggedin")
    if (a=="true"){
      router.navigate(['/'])
    }
  }

  ngOnInit(): void {
    

  }

  login() {
    this.loginRequestPayload.username = this.loginForm.get('username')?.value;
    this.loginRequestPayload.password = this.loginForm.get('password')?.value;
 console.log(this.loginRequestPayload.username,this.loginRequestPayload.password)
    this.authService.LogIn(this.loginRequestPayload).subscribe(data=>{
      console.log(data.email,data.username)
      localStorage.setItem("username",data.username)
      localStorage.setItem("email",data.email)
      localStorage.setItem("loggedin","true")
      location.reload()
      
    })
  }

}
