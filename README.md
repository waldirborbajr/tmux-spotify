<img src="https://github.com/user-attachments/assets/9fb07638-5907-4915-b9bf-1ca89255a93d" alt="drawing" style="width:100px;"/>

# Tmux Spotify Music Display

<p align="center">
  <img width="256" height="256" src="https://github.com/user-attachments/assets/c3771c8f-e066-4ed8-be91-3f73019849a2" />
</p>

This CLI application displays the currently playing music from Spotify on your tmux status bar.

## Prerequisites

- Go 1.16 or later
- tmux
- Spotify Premium account

## Installation

1. Clone this repository:
   ```
   git clone https://github.com/yourusername/tmux-spotify-cli.git
   cd tmux-spotify-cli
   ```

2. Create a Spotify application:
   - Go to https://developer.spotify.com/dashboard/
   - Create a new application
   - In your app settings, add `http://localhost:8080/callback` to the Redirect URIs
   - Note down the Client ID and Client Secret

3. Create a `.tmux-spotify-env` file in your home directory:
   ```
   echo "SPOTIFY_ID=your_client_id" > ~/.tmux-spotify-env
   echo "SPOTIFY_SECRET=your_client_secret" >> ~/.tmux-spotify-env
   echo "SPOTIFY_REDIRECT_URI=http://localhost:8080/callback" >> ~/.tmux-spotify-env
   ```
   Replace `your_client_id` and `your_client_secret` with the values from your Spotify app.

4. Build the application:
   ```
   make build
   ```

5. Add the following line to your `.tmux.conf`:
   ```
   set -g status-interval 5
   ```

## Usage

1. Run the application:
   ```
   ./tmux-spotify-cli
   ```

2. Follow the URL provided to authorize the application with your Spotify account.

3. The currently playing track will now appear in your tmux status bar.

## Troubleshooting

If you encounter an "INVALID_CLIENT: Invalid redirect URI" error:
1. Double-check that the redirect URI in your `.tmux-spotify-env` file matches exactly with what you've set in your Spotify app settings.
2. Ensure there are no trailing spaces or newlines in the `.tmux-spotify-env` file.
3. Verify that you've added the correct redirect URI to your Spotify app settings in the Spotify Developer Dashboard.

## Development

- Use `make dev` to build for development with race detection enabled.
- Use `make test` to run tests.
- Use `make build` to build for production.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
