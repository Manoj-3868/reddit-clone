import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './Components/home/home.component';
import { LoginComponent } from './Components/Auth/login/login.component';
import { SignupComponent } from './Components/Auth/signup/signup.component';
import { CreateSubredditComponent } from './Components/Subreddit/create-subreddit/create-subreddit.component';
import { ViewallSubredditComponent } from './Components/Subreddit/viewall-subreddit/viewall-subreddit.component';
import { CreatePostComponent } from './Components/Post/create-post/create-post.component';
import { ViewProfileComponent } from './Components/view-profile/view-profile.component';
import { ViewPostComponent } from './Components/Post/view-post/view-post.component';
import { AuthGuardGuard } from './Guard/auth-guard.guard';

const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'login', component: LoginComponent },
  { path: 'signup', component: SignupComponent },
  { path: 'create-subreddit', component: CreateSubredditComponent,canActivate:[AuthGuardGuard] },
  { path: 'list-subreddits', component: ViewallSubredditComponent},
  { path: 'create-post', component: CreatePostComponent,canActivate:[AuthGuardGuard]},
  { path: 'user-profile/:name', component: ViewProfileComponent,canActivate:[AuthGuardGuard]},
  { path: 'view-post/:id', component: ViewPostComponent,canActivate:[AuthGuardGuard] },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
