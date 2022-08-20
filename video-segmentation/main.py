import boto3

# get video duration
def get_video_duration(filename):
    """
    Get video duration
    """
    clip = VideoFileClip(filename)
    duration = clip.duration
    clip.close()
    return duration

# ffmpeg subclip video file
def ffmpeg_subclip_video_file(filename, t1, t2):
    """
    Ffmpeg subclip video file
    """
    subprocess.call(['ffmpeg', '-i', filename, '-ss', str(t1), '-to', str(t2), '-c', 'copy', '-y', filename.split('.')[0] + '_subclip.mp4'])
    return

#  video subclip
def video_subclip(filename, t1, t2):
    """
    Video clip
    """
    clip = VideoFileClip(filename)
    clip = clip.subclip(t1, t2)
    clip.write_videofile(filename.split('.')[0] + '_subclip.mp4')
    clip.close()
    return


# upload to s3  and delete local file
def upload_to_s3(filename, bucket_name):
    """
    Upload to s3
    """
    s3 = boto3.resource('s3')
    s3.meta.client.upload_file(filename, bucket_name, filename)
    os.remove(filename)
    return

# download from s3 
def download_from_s3(filename, bucket_name):
    """
    Download from s3
    """
    s3 = boto3.resource('s3')
    s3.Bucket(bucket_name).download_file(filename, filename)
    return


# print now  millisecond
def print_now_millisecond():
    """
    Print now millisecond
    """
    print(datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S.%f"))
    return

import math
for item in range(0, math.ceil(duration / 10)):
    print_now_millisecond()
    time.sleep(1)
    pass


# string concat
def string_concat(string1, string2):
    """
    String concat
    """
    return string1 + string2


# iterate list specific step size
def iterate_list_specific_step_size(list, step_size):
    """
    Iterate list specific step size
    """
    for i in range(0, len(list), step_size):
        yield list[i:i + step_size]
    return


# get video frame rate
def get_video_frame_rate(filename):
    """
    Get video frame rate
    """
    clip = VideoFileClip(filename)
    frame_rate = clip.fps
    clip.close()
    return frame_rate
