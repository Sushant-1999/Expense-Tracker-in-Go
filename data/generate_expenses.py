import random
from datetime import datetime, timedelta
import json
import os

# random.seed(42)

# Static values
users = ["john_doe", "jane_smith", "bob_jones", "alice_green", "sam_wilson"]
categories = ["GRS", "UTL", "RNT", "TRN", "DNG", "ENT", "HLC", "CLT", "EDU", "TRV", "PRC", "GFT", "HOM", "INS", "SAV"]

# expenses object with list
expenses = {"expenses": []}

# Generate expenses for a month before today
for day in range(1, 31):
    date = (datetime.now() - timedelta(days=day)).strftime("%Y-%m-%d")
    
    for _ in range(random.randint(1, len(users))):
        # Generate transaction values
        transaction_id = random.randint(1000, 10000)
        amount = round(random.uniform(1, 100), 2)
        category = random.choice(categories)
        user = random.choice(users)
        
        # Create expense object
        expense = {
            "transactionID": transaction_id,
            "amount": amount,
            "date": date,
            "category": category,
            "user": user
        }
        
        # Add to the expense list
        expenses["expenses"].append(expense)
        

# Save the generated expenses
if os.getcwd() != "data":
    os.chdir('data')

with open('./expenses.json', 'w') as f:
    json.dump(expenses, f, indent="\t")
