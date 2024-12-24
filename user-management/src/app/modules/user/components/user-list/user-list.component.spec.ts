import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UserListComponent } from './user-list.component';
import { UserService, User } from '../../../../core/services/user.service';
import { of } from 'rxjs';

const mockUsers: User[] = [
  { id: 1, user_name: 'John Doe', email: 'f4iMg@example.com' },
  { id: 2, user_name: 'Jane Smith', email: 'GZLbY@example.com' },
  { id: 3, user_name: 'Bob Johnson', email: 'Ewq0a@example.com' },
];

const userService = jasmine.createSpyObj('UserService', ['getUsers']);
describe('UserListComponent', () => {
  let component: UserListComponent;
  let fixture: ComponentFixture<UserListComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [UserListComponent]
    });
    fixture = TestBed.createComponent(UserListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should load users on initialization', () => {
    spyOn(userService, 'getUsers').and.returnValue(of(mockUsers));
    component.ngOnInit();
    expect(userService.getUsers).toHaveBeenCalled();
    expect(component.dataSource.data.length).toBe(mockUsers.length);
  });
  
});
