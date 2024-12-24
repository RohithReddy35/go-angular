import { Component, ChangeDetectionStrategy, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { UserFormComponent } from '../user-form/user-form.component';
import { MatSnackBar } from '@angular/material/snack-bar';
import { UserService, User } from 'src/app/core/services/user.service';
import { ConfirmDialogComponent } from 'src/app/shared/components/confirm-dialog/confirm-dialog.component';
import { MatTableDataSource } from '@angular/material/table';
import { MatPaginator } from '@angular/material/paginator';
import { MatSort } from '@angular/material/sort';
import { ViewChild } from '@angular/core';

@Component({
  selector: 'app-user-list',
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class UserListComponent implements OnInit {
  users: User[] = [];
  displayedColumns: string[] = ['id', 'user_name', 'email', 'actions'];

  constructor(private userService: UserService, private dialog: MatDialog, private snackBar: MatSnackBar) {}

  dataSource!: MatTableDataSource<User>;
  @ViewChild(MatPaginator) paginator!: MatPaginator;
  @ViewChild(MatSort) sort!: MatSort;
  ngOnInit(): void {
    this.loadUsers();
  }
  
  loadUsers(): void {
    this.userService.getUsers().subscribe((users) => {
      this.dataSource = new MatTableDataSource(users);
      this.dataSource.paginator = this.paginator;
      this.dataSource.sort = this.sort;
    });
  }

  openCreateDialog(): void {
    const dialogRef = this.dialog.open(UserFormComponent, {
      width: '400px',
      data: { isEdit: false }
    });
  
    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.loadUsers(); // Reload the user list
        this.snackBar.open('User created successfully!', 'Close', { duration: 3000 });
      }
    });
  }
  
  openEditDialog(user: User): void {
    const dialogRef = this.dialog.open(UserFormComponent, {
      width: '400px',
      data: { isEdit: true, user }
    });
  
    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.loadUsers(); // Reload the user list
        this.snackBar.open('User updated successfully!', 'Close', { duration: 3000 });
      }
    });
  }


confirmDelete(id: number): void {
  const dialogRef = this.dialog.open(ConfirmDialogComponent, {
    width: '300px'
  });

  dialogRef.afterClosed().subscribe((confirmed) => {
    if (confirmed) {
      this.userService.deleteUser(id).subscribe(() => {
        this.loadUsers(); // Reload the user list
        this.snackBar.open('User deleted successfully!', 'Close', { duration: 3000 });
      });
    }
  });
}

}
