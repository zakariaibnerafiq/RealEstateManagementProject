import { useEffect, useState } from "react";
import { Route, Routes } from "react-router-dom";
import "./App.css";

import Navbar from "./components/Navbar/Navbar";
import Aboutpage from "./pages/aboutpage/About";
import Homepage from "./pages/homepage/Home";
import Signinpage from "./pages/loginpage/Signin";
import Signuppage from "./pages/loginpage/Signup";
import Servicespage from "./pages/servicespage/Service";
import Propertiespage from "./pages/propertiespage/Properties";

// admin panel
import AdminPanel from "./Dashboard/Admin/Dashboard";
import AddProperty from "./Dashboard/Pages/Addproperty/AddProperty";
import AddUserPage from "./Dashboard/Pages/Adduser/AddUser";
import AllUser from "./Dashboard/Pages/Alluser/AllUser";
import UserProfile from "./Dashboard/Pages/UserProfile/UserProfile";
import UserVerification from "./Dashboard/Pages/UserVerification";
import VerifyUserPage from "./Dashboard/Pages/Verification/VerifyUser";
import VerifyProperty from "./Dashboard/Pages/Verification/VerifyProperty";
import Adminsaleproperty from "./Dashboard/Pages/Allproperty/Saleproperty";
import Adminrentproperty from "./Dashboard/Pages/Allproperty/Rentproperty";

// user dashboard Property Handler
import ActiveListings from "./Dashboard/Pages/Myproperty/ActiveListings";
import PendingListings from "./Dashboard/Pages/Myproperty/PendingListings";
import RentListings from "./Dashboard/Pages/Myproperty/RentListings";
import SoldListings from "./Dashboard/Pages/Myproperty/SoldListings";
import MyPurchasedListings from "./Dashboard/Pages/Myproperty/MyPurchasedListings";
import MyRentedListings from "./Dashboard/Pages/Myproperty/MyRentedListings";


import { useLocation } from "react-router-dom";
import { getID, setProfile } from "./utils/localstorage";

import { UserApi } from "./api/user";
import RentPage from "./pages/rentpage/Rentpage";
import SalePage from "./pages/salepage/Salepage";

function App() {
    const location = useLocation();
    const [loading, setLoading] = useState(true);
    const [isloggedin, setIsloggedin] = useState(false);
    useEffect(() => {
        const fetchUserProfile = async () => {
            setLoading(true);
            try {
                const userId = getID();
                if (userId) {
                    const { statusCode, data } = await UserApi.getProfile();
                    if (statusCode === 200) {
                        setProfile(data); // Ensure setProfile is properly defined
                    }
                }
            } catch (error) {
                console.error("Error fetching user profile", error);
            }
            setLoading(false);
        };
        fetchUserProfile();
    }, [isloggedin]);

    return (
        <>
            {!location.pathname.startsWith("/dashboard") && (
                <Navbar loading={loading} />
            )}
            <Routes>
                <Route path="/" element={<Homepage />} />
                <Route path="/buyproperty" element={<SalePage />} />
                <Route path="/rentproperty" element={<RentPage />} />
                <Route path="/buyproperty/:id" element={<Propertiespage />} />
                <Route path="/rentproperty/:id" element={<Propertiespage />} />
                <Route path="/about" element={<Aboutpage />} />
                <Route path="/services" element={<Servicespage />} />
                <Route
                    path="/signin"
                    element={<Signinpage loggedin={setIsloggedin} />}
                />
                <Route path="/signup" element={<Signuppage />} />

                <Route path="/dashboard" element={<AdminPanel loading={loading} />}>
                    <Route path="alluser" element={<AllUser />} />
                    <Route path="userprofile" element={<UserProfile />} />
                    <Route path="adduser" element={<AddUserPage />} />
                    <Route path="verifyuser" element={<VerifyUserPage />} />
                    <Route path="verifyproperty" element={<VerifyProperty />} />
                    <Route path="addproperty" element={<AddProperty />} />
                    <Route path="allsaleproperty" element={<Adminsaleproperty />} />
                    <Route path="allrentproperty" element={<Adminrentproperty />} />
                    <Route path="activelistings" element={<ActiveListings />} />
                    <Route path="pendinglistings" element={<PendingListings />} />
                    <Route path="rentproperties" element={<RentListings />} />
                    <Route path="soldproperties" element={<SoldListings />} />
                    <Route path="mypurchasedproperites" element={<MyPurchasedListings />} />
                    <Route path="myrentedproperties" element={<MyRentedListings />} />

                    <Route path="test" element={<UserVerification />} />
                </Route>
            </Routes>
        </>
    );
}

export default App;
