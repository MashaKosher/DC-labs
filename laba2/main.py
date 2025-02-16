from fastapi import Depends, FastAPI
from db_helper import db_helper
from contextlib import asynccontextmanager
from models import Base
from api_v1 import router as api_v1_router
from sqlalchemy.ext.asyncio import AsyncSession



from sqlalchemy import Column, Integer, String
from sqlalchemy.orm import declarative_base

# Base = declarative_base()

# class User(Base):
#     __tablename__ = 'users'

#     id = Column(Integer, primary_key=True, index=True)
#     name = Column(String, index=True)



@asynccontextmanager
async def lifespan(app: FastAPI):
    async with db_helper.engine.begin() as conn:
        await conn.run_sync(Base.metadata.create_all)
    yield

app = FastAPI(
    lifespan=lifespan,
    title = "Laba 2 service", 
)




# @app.post("/users/")
# async def create_user(name: str, 
#                       db: AsyncSession = Depends(db_helper.session_dependency)):
#     new_user = User(name=name)
#     db.add(new_user)
#     await db.commit()
#     await db.refresh(new_user)
#     return new_user

# @app.get("/users/")
# async def read_users(skip: int = 0, limit: int = 10, db: AsyncSession = Depends(db_helper.session_dependency)):
#     result = await db.execute("SELECT * FROM users LIMIT :limit OFFSET :skip", {"limit": limit, "skip": skip})
#     users = result.scalars().all()
#     return users


app.include_router(api_v1_router, prefix="/api")