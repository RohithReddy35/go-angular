import { Component, OnInit } from '@angular/core';

@Component({
    selector: 'app-user-list',
    templateUrl: './user-list.component.html',
    styleUrls: ['./user-list.component.scss']
})
export class UserListComponent implements OnInit {
    displayedColumns: string[] = ['id', 'user_name', 'email'];
    users = [
        { id: 1, user_name: 'John Doe', email: 'john@example.com' },
        { id: 2, user_name: 'Jane Smith', email: 'jane@example.com' }
    ];

    ngOnInit(): void {}
}
