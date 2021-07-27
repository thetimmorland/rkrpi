import uvicorn
from fastapi import FastAPI
from fastapi.responses import JSONResponse
from fastapi.middleware.cors import CORSMiddleware

from .database import Database

db = Database()
app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["http://rkr.timothy-morland.com"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


@app.get("/")
async def root(offset=0, limit=1000):
    msgs = db.read_msgs(offset, limit)
    return msgs


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8080)
