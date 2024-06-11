import getHandler from '@/handlers/get_handler';
import Toaster from '@/utils/toaster';
import React, { useEffect, useState } from 'react';
import { SERVER_ERROR } from '@/config/errors';
import Loader from '@/components/loader';
import { User } from '@/types';
import InfiniteScroll from 'react-infinite-scroll-component';
import UserComponent from '@/components/user';

const Users = () => {
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(true);
  const [hasMore, setHasMore] = useState(true);
  const [page, setPage] = useState(1);

  const getUsers = () => {
    setLoading(true);
    const URL = `/flags/users?page=${page}&limit=${10}`;
    getHandler(URL)
      .then(res => {
        if (res.statusCode === 200) {
          const addedUsers = [...users, ...(res.data.users || [])];
          if (addedUsers.length === users.length) setHasMore(false);
          setUsers(addedUsers);
          setPage(prev => prev + 1);
          setLoading(false);
        } else {
          if (res.data.message) Toaster.error(res.data.message, 'error_toaster');
          else {
            Toaster.error(SERVER_ERROR, 'error_toaster');
          }
        }
      })
      .catch(err => {
        Toaster.error(SERVER_ERROR, 'error_toaster');
      });
  };

  useEffect(() => {
    getUsers();
  }, []);

  return (
    <div className="w-full h-full p-4">
      {loading ? (
        <div className="w-[45vw] mx-auto max-lg:w-[85%] max-md:w-full">
          <Loader />
        </div>
      ) : (
        <InfiniteScroll
          className="w-[45vw] mx-auto max-lg:w-[85%] max-md:w-screen flex flex-col gap-2 max-lg:px-4 pb-base_padding"
          dataLength={users.length}
          next={getUsers}
          hasMore={hasMore}
          loader={<Loader />}
        >
          {users.length === 0 ? (
            <div>No Flagged Users.</div>
          ) : (
            users.map(user => <UserComponent key={user.id} user={user} setUsers={setUsers} />)
          )}
        </InfiniteScroll>
      )}
    </div>
  );
};

export default Users;
