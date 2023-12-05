import web
import urllib.parse
import subprocess, shlex

secret=xxxxx

# /Users/$(whoami)/Downloads/
download_dest_path=xxx

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
