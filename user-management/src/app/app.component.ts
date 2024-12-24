import { Component } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { UserFormComponent } from './modules/user/components/user-form/user-form.component';
import { UserService } from './core/services/user.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  constructor(
    private dialog: MatDialog,
    private snackBar: MatSnackBar,
    private userService: UserService
  ) {}

  openCreateDialog(): void {
    const dialogRef = this.dialog.open(UserFormComponent, {
      width: '400px',
      data: { isEdit: false }
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.snackBar.open('User created successfully!', 'Close', {
          duration: 3000
        });

        // Refresh the user list by triggering an event or directly reloading
        this.userService.refreshPage();
      }
    });
  }
}
