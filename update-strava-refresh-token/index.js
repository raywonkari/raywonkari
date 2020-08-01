// read strava refresh token
const fs = require('fs')
const value = fs.readFileSync('../strava_refresh_token', 'utf8')

// delete file, so we don't expose the token in the next PR step in GitHub Actions
fs.unlinkSync('../strava_refresh_token')

// Get GitHub Repo Public Key
const request = require('request');

// GitHub credentials. I am using the access token.
const token = process.env.MY_GITHUB_AUTH

// request details to get the public key of the repo
const options = {
	'method': 'GET',
	'url': 'https://api.github.com/repos/raywonkari/raywonkari/actions/secrets/public-key',
	'headers': {
		'Accept': 'application/vnd.github.v3+json',
		'Authorization': `${token}`,
		'User-Agent': 'nodejs-app-get',
	}
};

request(options, function (error, response) {
	if (error) console.log(error)
	const github = JSON.parse(response.body)

	// start secret encryption process
	// lib sodium
	const sodium = require('tweetsodium')

	// convert key and val to Uint8Array
	const messageBytes = Buffer.from(value)
	const keyBytes = Buffer.from(github.key, 'base64')

	// encrypt using Lib Sodium
	const encryptedBytes = sodium.seal(messageBytes, keyBytes)

	// base64 the encrypted secret
	const encrypted = Buffer.from(encryptedBytes).toString('base64')

	// request details for updating GH secret
	const newoptions = {
		'method': 'PUT',
		'url': 'https://api.github.com/repos/raywonkari/raywonkari/actions/secrets/STRAVA_REFRESH_TOKEN',
		'body': `{
          'encrypted_value': ${encrypted},
          'key_id': ${github.key_id}
        }`,
		'headers': {
			'accept': 'application/vnd.github.v3+json',
			'Authorization': `${token}`,
			'User-Agent': 'nodejs-app-put'
		},
	};

	request(newoptions, function (error, response) {
		if (error) console.log(error)
	});
});
