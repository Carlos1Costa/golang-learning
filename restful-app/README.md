# restful-app

This is a simple Go project to create a RESTful web server with CRUD operations for managing a list of music albums. The application uses a CSV file as the database and provides a web interface for interacting with the data.

## Requirements

- Go installed (version 1.24.2 or higher).
- A code editor like Visual Studio Code with the Go extension installed.

## Project Setup

1. Verify that Go is installed:
   ```bash
   go version
   ```

2. Clone this repository or navigate to the project directory:
   ```bash
   cd restful-app
   ```

3. Ensure the `albums.csv` file exists in the project directory. If it does not exist, the application will create an empty file when it runs.

## Running the Application

To start the server, run the following command from the `restful-app` directory:

```bash
go run main.go
```

The server will start on `http://localhost:8080`. You can access the application in your browser or use tools like `curl` to interact with it.

## Application Features

### CRUD Operations

The application implements the following CRUD operations:

1. **Create**: Add a new album using the "Add New Album" button on the main page. This opens a form where you can input the album's details (Name, Artist, Year).
2. **Read**: View all albums in a table on the main page (`/`). Each album is displayed with its ID, Name, Artist, and Year.
3. **Update**: Edit an existing album by clicking the "Edit" button next to it. This opens the same form used for creating albums, pre-filled with the album's details.
4. **Delete**: Delete an album by clicking the "Delete" button next to it. A confirmation prompt ensures that the deletion is intentional.

### CSV Database

- The application uses a CSV file (`albums.csv`) to store album data persistently.
- Each album is represented as a row in the CSV file with the following fields:
  - `ID`: A unique identifier for the album.
  - `Name`: The name of the album.
  - `Artist`: The name of the band or musician.
  - `Year`: The year the album was released.

### HTML Templates

- The application uses Go's `html/template` package to render HTML pages.
- Templates are located in the `templates/` directory:
  - `list.html`: Displays the table of albums and provides buttons for adding, editing, and deleting albums.
  - `form.html`: Provides a form for creating or editing an album.

### Mutex for Concurrency

- The application uses a `sync.Mutex` (`albumsMutex`) to ensure thread-safe access to the `albums` slice.
- The mutex is locked and unlocked around critical sections where the `albums` slice is modified or read.
- Functions like `saveAlbums` and handlers like `deleteAlbumHandler` and `saveAlbumHandler` carefully manage the mutex to avoid deadlocks.

### Logging

- The application logs important events, such as entering and exiting handlers, loading and saving albums, and errors.
- Logs are printed to the console to assist with debugging and monitoring.

## Code Structure

- **`main.go`**: The main application file containing the server setup, route handlers, and core logic.
- **`templates/`**: Contains the HTML templates for rendering the web interface.
- **`albums.csv`**: The CSV file used as the database for storing album data.

### Key Functions

1. **`loadAlbums`**:
   - Reads the `albums.csv` file and populates the `albums` slice.
   - Logs each album loaded from the file.

2. **`saveAlbums`**:
   - Writes the current state of the `albums` slice to the `albums.csv` file.
   - Ensures thread safety by locking the mutex during the write operation.

3. **Handlers**:
   - `listAlbumsHandler`: Renders the main page with the table of albums.
   - `newAlbumHandler`: Renders the form for creating a new album.
   - `editAlbumHandler`: Renders the form for editing an existing album.
   - `deleteAlbumHandler`: Deletes an album and updates the CSV file.
   - `saveAlbumHandler`: Handles form submissions for creating or updating albums.

## Example Usage

1. Start the server:
   ```bash
   go run main.go
   ```

2. Open your browser and navigate to `http://localhost:8080`.

3. Add a new album:
   - Click "Add New Album".
   - Fill in the form with the album's details.
   - Click "Save".

4. Edit an album:
   - Click "Edit" next to the album you want to modify.
   - Update the details in the form.
   - Click "Save".

5. Delete an album:
   - Click "Delete" next to the album you want to remove.
   - Confirm the deletion in the prompt.

6. View the updated list of albums on the main page.

## Notes

- The application automatically creates the `albums.csv` file if it does not exist.
- The server must be restarted to apply changes to the code or templates.

## Future Improvements

- Add pagination for the album list if the number of albums grows large.
- Implement search and filtering functionality.
- Add unit tests for the handlers and core functions.
- Improve error handling and user feedback in the web interface.

Feel free to contribute or suggest improvements!
