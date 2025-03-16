import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './Components/header/header.component';
import { LoginComponent } from './Components/Auth/login/login.component';
import { SignupComponent } from './Components/Auth/signup/signup.component';
import { HomeComponent } from './Components/home/home.component';
import { SideBarComponent } from './Components/Side-View/side-bar/side-bar.component';
import { SubredditSideBarComponent } from './Components/Side-View/subreddit-side-bar/subreddit-side-bar.component';
import { CreateSubredditComponent } from './Components/Subreddit/create-subreddit/create-subreddit.component';
import { ViewallSubredditComponent } from './Components/Subreddit/viewall-subreddit/viewall-subreddit.component';
import { NgxUiLoaderHttpModule, NgxUiLoaderModule } from 'ngx-ui-loader';
import { ViewAllPostComponent } from './Components/Post/view-all-post/view-all-post.component';
import { CreatePostComponent } from './Components/Post/create-post/create-post.component';
import { ViewProfileComponent } from './Components/view-profile/view-profile.component';
import { ViewPostComponent } from './Components/Post/view-post/view-post.component';
import { ManageVoteComponent } from './Components/Vote/manage-vote/manage-vote.component';


@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    LoginComponent,
    SignupComponent,
    HomeComponent,
    SideBarComponent,
    SubredditSideBarComponent,
    CreateSubredditComponent,
    ViewallSubredditComponent,
    ViewAllPostComponent,
    CreatePostComponent,
    ViewProfileComponent,
    ViewPostComponent,
    ManageVoteComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    ReactiveFormsModule,
    HttpClientModule,
    NgxUiLoaderModule,
    NgxUiLoaderHttpModule.forRoot({
      showForeground:true,
    }),
    
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
