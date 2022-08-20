# import 
import subprocess
import argparse

# ffmpeg subclip video file by i frame number
def ffmpeg_subclip_video_file_by_i_frame_number(filename, i_frame_number):
    """
    Ffmpeg subclip video file by i frame number
    """
    subprocess.call(['ffmpeg', '-i', filename, '-vf', 'select=gt(n\,' + str(i_frame_number) + ')', '-y', filename.split('.')[0] + '_subclip.mp4'])
    return


# ffmpeg subclip video file by i frame
def ffmpeg_subclip_video_file_by_i_frame(filename, i_frame):
    """
    Ffmpeg subclip video file by i frame
    """
    subprocess.call(['ffmpeg', '-i', filename, '-vf', 'select=gt(n\,' + str(i_frame) + ')', '-y', filename.split('.')[0] + '_subclip.mp4'])
    return

# ffmpeg get video i frame number
def ffmpeg_get_video_i_frame_number(filename):
    """
    Ffmpeg get video i frame number
    """
    return subprocess.call(['ffmpeg', '-i', filename, '-vf', 'showinfo', '-frames:v', '1', '-an', '-f', 'null', '-'])


# ffmpeg get video frame info by i frame number
def ffmpeg_get_video_frame_info_by_i_frame_number(filename):
    """
    Ffmpeg get video frame info by i frame number
    """
    return subprocess.call(['ffmpeg', '-i', filename, '-vf', 'showinfo', '-frames:v', '1', '-an', '-f', 'null', '-'])

## ffmpeg get video all frame info
def ffmpeg_get_video_all_frame_info_and_print(filename):
    """
    Ffmpeg get video all frame info and print
    """
    return subprocess.call(['ffmpeg', '-i', filename, '-vf', 'showinfo', '-frames:v', '-1', '-an', '-f', 'null', '-'])

# ffmpeg -i test.mp4 -vf "select='eq(pict_type,I)'" -vsync vfr out-%02d.jpeg
# https://jdhao.github.io/2021/12/25/ffmpeg-extract-key-frame-video/


# 1) Fast Way to Cut / Trim Without Re-encoding (using Copy and Input Seeking)
# 2) Cut/Trim using Output Seeking Without Re-encoding
# https://ottverse.com/trim-cut-video-using-start-endtime-reencoding-ffmpeg/#Fast_Way_to_Cut_Trim_Without_Re-encoding_using_Copy_and_Input_Seeking

'''
Since the seeking operation jumps between I-frames, it is not going to accurately stop on the frame (or time) that you requested. It will search for the nearest I-frame and start the copy operation from that point.
'''
# main function

if __name__ == '__main__':
    # parse arguments
    parser = argparse.ArgumentParser()
    parser.add_argument('-f', '--filename', help='filename', required=True)
    parser.add_argument('-i', '--i_frame_number', help='i frame number', required=True)
    args = parser.parse_args()

    # ffmpeg get i frame number
    #  print(str(ffmpeg_get_video_i_frame_number(args.filename)) + '\n')
    ffmpeg_get_video_frame_info_by_i_frame_number(args.filename)
    #  ffmpeg_subclip_video_file_by_i_frame_number(args.filename, args.i_frame_number)
    #  ffmpeg_subclip_video_file_by_i_frame(args.filename, args.i_frame_number)
    #  ffmpeg_get_video_all_frame_info_and_print(args.filename)


