import { Component, Inject, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { User, UserService } from 'src/app/core/services/user.service';
import { debounceTime as rxjsDebounceTime } from 'rxjs/operators';

@Component({
  selector: 'app-user-form',
  templateUrl: './user-form.component.html',
  styleUrls: ['./user-form.component.scss']
})
export class UserFormComponent implements OnInit {
  userForm!: FormGroup;

  constructor(
    private fb: FormBuilder,
    private userService: UserService,
    private dialogRef: MatDialogRef<UserFormComponent>,
    @Inject(MAT_DIALOG_DATA) public data: { isEdit: boolean; user?: User }
  ) {}

  ngOnInit(): void {
    this.userForm = this.fb.group({
      user_name: [this.data.user?.user_name || '', Validators.required],
      email: [
        this.data.user?.email || '',
        [Validators.required, Validators.email]
      ]
    });
    this.userForm.get('email')?.valueChanges.pipe(debounceTime(300)).subscribe((value) => {
      console.log('Debounced email value:', value);
    });    
  }

  onSubmit(): void {
    if (this.data.isEdit) {
      console.log('Updating user:', this.userForm.value);
      this.userService
        .updateUser(this.data.user!.id!, this.userForm.value)
        .subscribe(() => {
          this.dialogRef.close(true); // Signal success
        });
    } else {
      this.userService.createUser(this.userForm.value).subscribe(() => {
        this.dialogRef.close(true); // Signal success
      });
    }
  }
}
function debounceTime<T>(dueTime: number): import("rxjs").OperatorFunction<T, T> {
  return rxjsDebounceTime(dueTime);
}

