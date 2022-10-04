

state = False



visits = list()

for i in range(100):
    visits.append(i)


blub_list = []
prev_blub_list = []
blub_state = dict()
for num in range(100):
    
    blub_list=visits[num:100:(num+1)]
    prev_blub_list = blub_list.copy()
    if num ==1:
        continue
    
    for num_list in blub_list:
        if num%2 == 0:
            prev_blub_list.remove(num_list)
        else:
            prev_blub_list.append(num_list)

print(prev_blub_list)  







