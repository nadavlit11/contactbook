run docker compose up to start the application

create a user via /users/register/:name API

to create contacts for this users use /contacts/add
send a post request with the following body:
{
    "first_name": "contact name",
    "last_name": "contact name",
    "phone": "contact phone",
    "address": "contact address"
}
and user_id in the header

to get all contacts for a user use /contacts/:page
send a get request with the user_id in the header and the page number in the url

to search for a contact use /contacts/search
send a get request with the user_id in the header and the search query in the url

to edit a contact use /contacts/edit
send a put request with the user_id in the header and the contact_id in the url and the new contact details in the body

to delete a contact use /contacts/delete
send a delete request with the user_id in the header and the contact_id in the url

to load user's contacts to cache use /users/login/:user_id
send a post request with the user_id in the url