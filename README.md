# tmux-spotify-cli

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

2. Build the application:
   ```
   make build
   ```

3. Create a Spotify application:
   - Go to https://developer.spotify.com/dashboard/
   - Create a new application
   - Set the redirect URI to `http://localhost:8080/callback`
   - Note down the Client ID and Client Secret

4. Create a `.tmux-spotify-env` file in your home directory:
   ```
   echo "SPOTIFY_ID=your_client_id" > ~/.tmux-spotify-env
   echo "SPOTIFY_SECRET=your_client_secret" >> ~/.tmux-spotify-env
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

## Development

- Use `make dev` to build for development with race detection enabled.
- Use `make test` to run tests.
- Use `make build` to build for production.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
