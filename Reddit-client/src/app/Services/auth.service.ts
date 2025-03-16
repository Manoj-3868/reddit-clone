import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, Observable, retry, throwError } from 'rxjs';
import { Login } from '../Models/login';
import { Signup } from '../Models/signup';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  httpHeader={
    headers:new HttpHeaders({
      'Content-Type':'application/json'
    }
    )
  }
  baseUrl:string="http://localhost:3000"
  constructor(private httpClient:HttpClient) { }


  SignUp(signup:Signup):Observable<Signup>{
    return this.httpClient.post<Signup>(this.baseUrl+'/signup',JSON.stringify(signup),this.httpHeader)
    .pipe(
      retry(1),
      catchError(this.httpError)
    );
  }

  LogIn(login:Login):Observable<Signup>{
    return this.httpClient.post<Signup>(this.baseUrl+'/login',JSON.stringify(login),this.httpHeader)
    .pipe(
      retry(1),
      catchError(this.httpError)
    );
  }

  httpError(error:HttpErrorResponse){
    let msg='';
    if(error.error instanceof ErrorEvent){
      msg=error.error.message;
    }
    else{
      msg=`Error Code:${error.status}\nMessafe:${error.message}`;
    }
    console.log(msg);
    return throwError(msg);
  }

}
