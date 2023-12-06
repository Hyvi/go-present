import web
import urllib.parse
import subprocess, shlex
import os

# Get From Env MY_DOWNLOAD_SECRET
secret=os.environ["MY_DOWNLOAD_SECRET"]
# Check if secret exists
if not secret:
    print("MY_DOWNLOAD_SECRET not exists")
    exit(1)

# Get /Users/$(whoami)/Downloads/
download_dest_path= os.path.join(os.path.expanduser('~'), 'Downloads')
# Check if download_dest_path exists
if not os.path.exists(download_dest_path):
    print("download_dest_path not exists")
    exit(1)

urls = (
    '/(.*)', 'hello'
)
app = web.application(urls, globals())

class hello:
    def GET(self, name):
        data = web.input()

        if 'file' not in data or 'url' not in data or 'secret' not in data:
            return 'server param err' 
        if data['secret'] != secret:
            return 'server err'

        # Check if contains "
        if '"' in data.file or '"' in data.url:
            return 'server err'

        command = 'you-get -x 127.0.0.1:7890 -o %s --output-filename "%s" "%s"' % (download_dest_path, urllib.parse.unquote(data.file), urllib.parse.unquote(data.url))
        if 'noproxy' in data:
            command = 'you-get -o %s --output-filename "%s" "%s"' % (download_dest_path, urllib.parse.unquote(data.file), urllib.parse.unquote(data.url))
        if data.file == 'bt':
            command = '/usr/local/opt/transmission-cli/bin/transmission-remote -a "%s"' % urllib.parse.unquote(data.url)
        print(command)
        subprocess.Popen(shlex.split(command), stdout=subprocess.PIPE, stderr=subprocess.PIPE)
         
        return 'Hello' + '!'

if __name__ == "__main__":
    app.run()
