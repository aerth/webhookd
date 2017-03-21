# webhookd

### http server for handling webhooks

Place shell scripts in "hooks" directory

Example:

./hooks/
./hooks/one
./hooks/two

"one" will be triggered with:
http://example.org/any/dir/one
or http://example.org/one

for easy fast use with reverse proxy or on its own

returns 404 for everything else
