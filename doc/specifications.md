# Housing research app

This application is a "one-shot" web application for searching a new house.

It will be used in our local network by me, Willow, and my family.
There is no need for authentication, users management, etc.

## Features

- Display a list of houses we are interested in.
- For each house, include the following information:
  - Creation date (automatic field)
  - Last update date (automatic field)
  - Title (mandatory, one-line text)
  - Publication URLs (optional, multiple URLs are allowed)
  - Publication date (mandatory for each publication URL, date with a calendar picker)
  - City (mandatory, selected in a pre-defined list)
  - Address (optional, free text)
  - Price, in euros (mandatory, integer number)
  - Surface (mandatory, integer number, in square meters)
  - Number of rooms (mandatory, integer number)
  - Number of bedrooms (mandatory, integer number)
  - Number of bathrooms (mandatory, integer number)
  - Number of floors (mandatory, integer number)
  - Year of construction (optional, integer number)
  - Type of house (house or apartment)
  - Total surface of the land (optional, integer number, in square meters)
  - Garage (optional, binary)
  - Number of outdoor parking spaces (optional, integer number)
  - Photos (optional, multiple files are allowed, one must be marked as the manually selected main photo)
  - Notes (optional, free text)
  - Other attached files (optional, multiple files are allowed)

## User interface

The user interface will be a web interface, composed of the following pages:

- Main page: summary of houses presented in a sortable table
- House details page: detailed view of a house
- Add new house page: form to add a new house
- Edit house page: form to edit an existing house
- Delete house page: confirmation to delete an existing house
- Modify cities page: form to modify the list of cities (once a city is used by at least one house, it must not be allowed to delete it)

The main color of the interface must be purple.
The text must be written in french, in the feminine form if needed, because the users are only women.
The URLs parts must be written in french.

All pages should have a menu fixed on the left side, to navigate between the different pages, with the following entries:

- Summary
- Add new house
- Modify cities
- Direct link to each house (thumbnail of the main photo and title of the house)

## Technical details

- Photos and attachments will be stored in subdirectories of an `uploads` directory, each subdirectory named after its house id, which will only be created on first run if it does not already exist. These files will be searched for directly on the filesystem, without any entry or table in the database. Only the main photo selection should appear in the database.
- Optional fields are stored in the database as the zero value of the type, not as nullable fields.
