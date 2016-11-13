# testing json comparison
# NEW reading in multiple JSON files from a sourceFolder
# AND reading in after making changes all JSON files from a targerFolder

#read in two json files and compare to see if any changes made
#print out any changes in files

import glob
import os, json
import sys


#sourcefolder = '/Users/arnoldas/Desktop/Fall 2016/hackRPI/sourcefolder/'
sourcefolder = directory_name=sys.argv[1]
#targetfolder = '/Users/arnoldas/Desktop/Fall 2016/hackRPI/targetfolder/'
targetfolder = directory_name=sys.argv[2]
#targetfile = './sourcefolder/'
#outputfilenameprefix = 'test1'




contentSource = []
contentTarget = []
#json_dir_name = "/path/to/json/dir"




#for filesystem before downloading anything or possibility of damage
json_filesSource = [pos_json for pos_json in os.listdir(sourcefolder) if pos_json.endswith('.json')]
print json_filesSource

for js in json_filesSource:
    with open(os.path.join(sourcefolder, js)) as json_file:
        contentSource.append(json.load(json_file))
        #print json.load(json_file)
        #print contentSource

#for filesystem after changes are made
json_filesTarget = [pos_json for pos_json in os.listdir(targetfolder) if pos_json.endswith('.json')]
print json_filesTarget

for js in json_filesTarget:
    with open(os.path.join(targetfolder, js)) as json_file:
        contentTarget.append(json.load(json_file))
        #print json.load(json_file)
        #print contentTarget

#print contentSource
# dictionaries are unhashable, let's convert to strings for sorting
sorted_1 = sorted([x for x in contentSource])
sorted_2 = sorted([x for x in contentTarget])

#print sorted_1
#UNCOMMENT THIS IF YOU WANT TO TEST THE CHANGE DETECTION ALGORITHM
'''
sorted_1[0][0]["FullPath"] = "e0w"
sorted_1[0][1]["PermissionNumber"] = 150

print sorted_1[0][0], sorted_2[0][0], (sorted_1[0][0] == sorted_2[0][0]), "\n\n"
'''

'''
print "\n\nprinting sorted_1 : from contentsource:\n\n\n"
for x in sorted_1[0]:
    print x
print "\n\nprinting sorted_2 : from contenttarget:\n\n\n"
for x in sorted_2[0]:
    print x
'''

'''
for x in range(0,len(sorted_1[0])-1):
    print x, " ", sorted_1[0][x]
    if sorted_1[0][x] in sorted_2[0]:
        sorted_2[0].remove(sorted_1[0][x])
        print sorted_2[0]
'''
# changes in 2 that are missing from 1
sorted_3 = [x for x in sorted_2[0] if x not in sorted_1[0]]
print sorted_3

# stuff from 1 that's changed or missing from 2
sorted_4 = [x for x in sorted_1[0] if x not in sorted_2[0]]

print sorted_3, "\n"
print sorted_4, "\n"

# With the above two arrays, generate a changelog

if (sorted_3 == []) and (sorted_4 == []):
    print "No changes found."

else:
    changelist = []
    findflag = False

    for x in sorted_3:
        for y in sorted_4:
            if x['FullPath'] == y['FullPath']:
                findflag = True
                if x['ContentChanged'] != y ['ContentChanged']:
                    changelist.append("Content changed: " + x['FullPath'])
                if x['LastAccessTime'] != y ['LastAccessTime']:
                    changelist.append("File accessed: " + x['FullPath'] + " @ " + str(y['LastAccessTime']))
                if x['PermissionNumber'] != y ['PermissionNumber']:
                    changelist.append("Permissions changed: " + x['FullPath'] + " from " + str(x['PermissionNumber']) + " to " + str(y['PermissionNumber']))
                break

        if (not findflag):
            changelist.append("File deleted or renamed: " + x['FullPath'])
        findflag = False

    for x in sorted_4:
        for y in sorted_3:
            if x['FullPath'] == y['FullPath']:
                findflag = True
                break

        if (not findflag):
            changelist.append("File added or renamed: " + x['FullPath'])
        findflag = False

    for s in changelist:
        print s




'''
# in case the dictionaries are all unique or you don't care about duplicities,
# sets should be faster than sorting
set_1 = set(repr(x) for x in contentTarget)
set_2 = set(repr(x) for x in contentTarget)
if set_1 == set_2:
    print "\neverything is safe and no changes found\n"
else:
    #add code for when cahnges are found
    print "\nchanges found here: \n"
print "Prints true if filesystem doesnt have changes and youre safe\n" \
      "False printed if changes found:\n" \
      "Result: "
print(set_1 == set_2)
'''
