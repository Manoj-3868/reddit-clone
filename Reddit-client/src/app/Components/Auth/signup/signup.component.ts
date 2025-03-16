import { Component, OnInit } from '@angular/core';
import { Signup } from 'src/app/Models/signup';
import { Router } from '@angular/router';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { AuthService } from 'src/app/Services/auth.service';
@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.scss']
})
export class SignupComponent implements OnInit {

  signupRequestPayload: Signup;
  signupForm: FormGroup;

  constructor(private authService:AuthService,private router:Router) {
    this.signupRequestPayload = {
      username: '',
      email: '',
      password: ''
    };
    this.signupForm = new FormGroup({
      username: new FormControl('', Validators.required),
      email: new FormControl('', [Validators.required, Validators.email]),
      password: new FormControl('', Validators.required),
    });
  }

  ngOnInit() {
   
  }

  signup() {
    this.signupRequestPayload.email = this.signupForm.get('email')?.value
    this.signupRequestPayload.username = this.signupForm.get('username')?.value;
    this.signupRequestPayload.password = this.signupForm.get('password')?.value;

    console.log(this.signupRequestPayload.email,this.signupRequestPayload.username,this.signupRequestPayload.password)
    this.authService.SignUp(this.signupRequestPayload).subscribe(data=>{
      console.log(data)
      this.router.navigate(['/login'])
    })
  }
}
