import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewallSubredditComponent } from './viewall-subreddit.component';

describe('ViewallSubredditComponent', () => {
  let component: ViewallSubredditComponent;
  let fixture: ComponentFixture<ViewallSubredditComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ViewallSubredditComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ViewallSubredditComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
