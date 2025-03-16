import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ManageVoteComponent } from './manage-vote.component';

describe('ManageVoteComponent', () => {
  let component: ManageVoteComponent;
  let fixture: ComponentFixture<ManageVoteComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ManageVoteComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ManageVoteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
